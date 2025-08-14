package com.joinself.app.academy

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
import com.joinself.sdk.ui.integrateUIFlows
import com.joinself.sdk.ui.openRegistrationFlow
import com.joinself.ui.theme.SelfModifier
import kotlinx.coroutines.launch
import java.io.File

class MainActivity : ComponentActivity() {
    val LOGTAG = "SelfSDK"
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        enableEdgeToEdge()

        println("🆕 New Account Creation Example")
        println("===============================")

        // Step 1: SDK initialization
        SelfSDK.initialize(applicationContext,
            log = { Log.d(LOGTAG, it) }
        )

        // Step 2: Account Storage Check
        Log.d(LOGTAG, "🔍 Checking for account storage path...")
        val storagePath = File(applicationContext.filesDir.absolutePath + "/new_account")
        if (!storagePath.exists()) storagePath.mkdirs()

        // Step 3: Account Initialization
        Log.d(LOGTAG,"🔧 Initialize Self account...")
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
                    Log.d(LOGTAG, "✅ Connected to Self network")
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

                    // Step 4: Account Registration Status
                    Log.d(LOGTAG,"🔧 Get registration status...")
                    var isRegistered by remember { mutableStateOf(account.registered()) }

                    NavHost(navController = navController, startDestination = "main", modifier = Modifier.padding(innerPadding)) {

                        // Step 3: UI Flow Integration
                        Log.d(LOGTAG,"🔧 Integrate Self-UI flows")
                        SelfSDK.integrateUIFlows(this, navController, selfModifier)

                        composable("main") {
                            Column(
                                verticalArrangement = Arrangement.spacedBy(10.dp),
                                horizontalAlignment = Alignment.CenterHorizontally,
                                modifier = Modifier.padding(start = 8.dp, end = 8.dp).fillMaxWidth()
                            ) {
                                Text(modifier = Modifier.padding(top = 40.dp), text = "Registered: $isRegistered")
                                Button(modifier = Modifier.padding(top = 20.dp),
                                    onClick = {
                                        Log.d(LOGTAG,"🔧 Open registration flow...")
                                        coroutineScope.launch {

                                            // Step 5: Account Registration Flow
                                            account.openRegistrationFlow { isSuccess, error ->
                                                // Step 6: Account Registered
                                                isRegistered = isSuccess
                                                Log.d(LOGTAG, "✅ Registration flow finished")
                                                Log.d(LOGTAG, "✅ New Self account ready!")
                                            }
                                        }
                                    },
                                    enabled = !isRegistered
                                ) {
                                    Text(text = "Create Account")
                                }
                            }
                        }
                    }
                }
            }
        }
    }
}