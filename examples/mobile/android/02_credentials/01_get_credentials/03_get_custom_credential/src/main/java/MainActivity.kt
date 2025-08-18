package com.joinself.app.academy

import Common
import android.os.Bundle
import android.util.Log
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.activity.enableEdgeToEdge
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.material3.Button
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Scaffold
import androidx.compose.material3.Text
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.rememberCoroutineScope
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.unit.dp
import androidx.compose.ui.window.Dialog
import androidx.compose.ui.window.DialogProperties
import androidx.navigation.compose.NavHost
import androidx.navigation.compose.composable
import androidx.navigation.compose.rememberNavController
import com.joinself.common.Environment
import com.joinself.sdk.SelfSDK
import com.joinself.sdk.models.Account
import com.joinself.sdk.models.CredentialMessage
import com.joinself.sdk.models.Message
import com.joinself.sdk.models.PublicKey
import com.joinself.sdk.ui.DisplayRequestUI
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
        val storagePath = File(applicationContext.filesDir.absolutePath + "/get_credentials")
        if (!storagePath.exists()) storagePath.mkdirs()

        var account: Account? = null

        setContent {
            MaterialTheme {
                Scaffold(
                    modifier = Modifier.fillMaxSize(),
                    containerColor = Color.White
                ) { innerPadding ->
                    val coroutineScope = rememberCoroutineScope()
                    val navController = rememberNavController()
                    val selfModifier = SelfModifier.sdk()

                    var isRegistered by remember { mutableStateOf(false) }
                    var groupAddress by remember { mutableStateOf<PublicKey?>(null) }
                    var statusText by remember { mutableStateOf("") }

                    var requestMessage by remember { mutableStateOf<CredentialMessage?>(null) }

                    LaunchedEffect(true) {
                        account = Account.Builder()
                            .setContext(applicationContext)
                            .setEnvironment(Environment.production)
                            .setSandbox(true)
                            .setStoragePath(storagePath.absolutePath)
                            .setCallbacks(object : Account.Callbacks {
                                override fun onMessage(message: Message) {
                                    Log.d(LOGTAG, "onMessage: ${message.id()}")
                                    if (message is CredentialMessage) requestMessage = message
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

                        isRegistered = account.registered()
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
                                            account?.openRegistrationFlow { isSuccess, error ->
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
                                            account?.openQRCodeFlow(
                                                onFinish = { qrCode, discoverData ->
                                                    coroutineScope.launch(Dispatchers.IO) {
                                                        groupAddress = Common.connect(account, qrCode)
                                                        statusText = if (groupAddress != null) "Server Connected!!" else "Failed to connect to Server!!"
                                                    }
                                                },
                                                onExit = {}
                                            )
                                        }
                                    },
                                    enabled = isRegistered,
                                ) {
                                    Text(text = "Scan QRCode")
                                }

                                Button(
                                    onClick = {
                                        Common.notifyServerForRequest( account ?:return@Button, groupAddress ?:return@Button,"REQUEST_GET_CUSTOM_CREDENTIAL")
                                    },
                                    enabled = isRegistered
                                ) {
                                    Text(text = "Get Credentials")
                                }

                                Text(text = statusText)

                                // display credential message with buttons to confirm or reject storing credentials
                                if (requestMessage != null) {
                                    Dialog(
                                        onDismissRequest = { },
                                        properties = DialogProperties(dismissOnBackPress = false, dismissOnClickOutside = false, usePlatformDefaultWidth = false),
                                    ) {
                                        account?.DisplayRequestUI(selfModifier, requestMessage ?: return@Dialog, onFinish = { isSent, status ->
                                            requestMessage = null
                                        })
                                    }
                                }
                            }
                        }
                    }
                }
            }
        }
    }
}