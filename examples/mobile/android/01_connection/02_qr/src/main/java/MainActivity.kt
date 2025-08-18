package com.joinself.app.academy

import android.os.Bundle
import android.util.Log
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.activity.enableEdgeToEdge
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.PaddingValues
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.width
import androidx.compose.material3.Button
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Scaffold
import androidx.compose.material3.Text
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.rememberCoroutineScope
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.unit.dp
import androidx.navigation.compose.NavHost
import androidx.navigation.compose.composable
import androidx.navigation.compose.rememberNavController
import com.joinself.common.Environment
import com.joinself.sdk.SelfSDK
import com.joinself.sdk.models.Account
import com.joinself.sdk.models.Message
import com.joinself.sdk.models.PublicKey
import com.joinself.sdk.ui.integrateUIFlows
import com.joinself.sdk.ui.openQRCodeFlow
import com.joinself.sdk.ui.openRegistrationFlow
import com.joinself.ui.theme.SelfModifier
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.launch
import java.io.File

class MainActivity : ComponentActivity() {
    val LOGTAG = "SelfSDK"
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        enableEdgeToEdge()

        SelfSDK.initialize(applicationContext,
            log = { Log.d(LOGTAG, it) }
        )

        // the sdk will store data in this directory, make sure it exists.
        val storagePath = File(applicationContext.filesDir.absolutePath + "/connection_qr")
        if (!storagePath.exists()) storagePath.mkdirs()

        val account = Account.Builder()
            .setContext(applicationContext)
            .setEnvironment(Environment.production)
            .setSandbox(true)
            .setStoragePath(storagePath.absolutePath)
            .setCallbacks(object : Account.Callbacks {
                override fun onMessage(message: Message) {
                    Log.d(LOGTAG, "onMessage: ${message.id()}")
                }
                override fun onConnect() {
                    Log.d(LOGTAG, "onConnect")
                }
                override fun onDisconnect(errorMessage: String?) {
                    Log.d(LOGTAG, "onDisconnect: $errorMessage")
                }
                override fun onAcknowledgement(id: String) {
                    Log.d(LOGTAG, "onAcknowledgement: $id")
                }
                override fun onError(id: String, errorMessage: String?) {
                    Log.d(LOGTAG, "onError: $errorMessage")
                }
            })
            .build()

        setContent {
            MaterialTheme {
                Scaffold(
                    modifier = Modifier.fillMaxSize(),
                    containerColor = Color.White
                ) { innerPadding ->
                    val coroutineScope = rememberCoroutineScope()
                    val navController = rememberNavController()
                    val selfModifier = SelfModifier.sdk()

                    var isRegistered by remember { mutableStateOf(account.registered()) }
                    var statusText by remember { mutableStateOf("") }

                    // connect with server by an inbox address, a group address is returned.
                    fun connect(qrCode: ByteArray) {
                        statusText = ""
                        coroutineScope.launch(Dispatchers.IO) {
                            try {
                                val groupAddress = account.connectWith(qrCode)
                                statusText = if (groupAddress != null) "Connected!!" else "Failed to connect!!"
                            } catch (ex: Exception) {
                                Log.e("Self", ex.message, ex)
                                statusText = "Failed to connect!!\n${ex.message}"
                            }
                        }
                    }

                    NavHost(navController = navController, startDestination = "main", modifier = Modifier.padding(innerPadding)) {
                        SelfSDK.integrateUIFlows(this, navController, selfModifier)

                        composable("main") {
                            Column(
                                verticalArrangement = Arrangement.spacedBy(20.dp),
                                horizontalAlignment = Alignment.CenterHorizontally,
                                modifier = Modifier.padding(start = 8.dp, end = 8.dp).fillMaxWidth()
                            ) {
                                Text(modifier = Modifier.padding(top = 40.dp), text = "Registered: $isRegistered")
                                Button(
                                    onClick = {
                                        coroutineScope.launch {
                                            // open registration flow to create an account
                                            account.openRegistrationFlow { isSuccess, error ->
                                                isRegistered = isSuccess
                                            }
                                        }
                                    },
                                    enabled = !isRegistered
                                ) {
                                    Text(text = "Create Account")
                                }

                                Button(
                                    onClick = {
                                        coroutineScope.launch {
                                            account.openQRCodeFlow(
                                                onFinish = { qrCode, discoverData ->
                                                    connect(qrCode)
                                                },
                                                onExit = {}
                                            )
                                        }
                                    },
                                    enabled = isRegistered,
                                ) {
                                    Text(text = "Scan QRCode")
                                }
                                Text(text = statusText)
                            }
                        }
                    }
                }
            }
        }
    }
}