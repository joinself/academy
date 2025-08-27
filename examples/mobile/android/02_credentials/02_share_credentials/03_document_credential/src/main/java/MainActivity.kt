package com.joinself.app.academy

import Common
import android.os.Bundle
import android.util.Log
import android.widget.Toast
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
import com.joinself.sdk.models.CredentialRequest
import com.joinself.sdk.models.Message
import com.joinself.sdk.models.PublicKey
import com.joinself.sdk.ui.DisplayRequestUI
import com.joinself.sdk.ui.integrateUIFlows
import com.joinself.sdk.ui.openDocumentVerificationFlow
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

        // Step 1: SDK initialization
        Log.d(LOGTAG, "🔧 Initialize Self SDK...")
        SelfSDK.initialize(applicationContext,
            log = { Log.d(LOGTAG, it) }
        )

        // the sdk will store data in this directory, make sure it exists.
        val storagePath = File(applicationContext.filesDir.absolutePath + "/share_custom_credential")
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
                    var isDocumentVerified by remember { mutableStateOf(false) }
                    var requestMessage by remember { mutableStateOf<CredentialRequest?>(null) }


                    LaunchedEffect(true) {
                        // Step 2: Account Initialization and UI Flow Integration
                        account = Account.Builder()
                            .setContext(applicationContext)
                            .setEnvironment(Environment.production)
                            .setSandbox(true)
                            .setStoragePath(storagePath.absolutePath)
                            .setCallbacks(object : Account.Callbacks {
                                override fun onMessage(message: Message) {
                                    Log.d(LOGTAG, "onMessage: ${message.id()}")
                                    // check if it is a liveness request
                                    if (message is CredentialRequest) {
                                        Log.d(LOGTAG, "✅ Received request message")
                                        requestMessage = message
                                    }
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

                        isRegistered = account.registered()
                    }

                    NavHost(navController = navController, startDestination = "main", modifier = Modifier.padding(innerPadding)) {
                        // Step 2: Account Initialization and UI Flow Integration
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
                                            // Step 3: Account Registration
                                            Log.d(LOGTAG, "🔧 Open registration flow...")
                                            account?.openRegistrationFlow { isSuccess, error ->
                                                isRegistered = isSuccess
                                                Log.d(LOGTAG, "✅ Registration successfully!")
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
                                            // Step 4: Scan QRCode
                                            Log.d(LOGTAG, "🔧 Open QRCode flow...")
                                            account?.openQRCodeFlow(
                                                onFinish = { qrCode, discoverData ->
                                                    coroutineScope.launch(Dispatchers.IO) {
                                                        groupAddress = Common.connect(account, qrCode)
                                                        statusText = if (groupAddress != null) "Server Connected!!" else "Failed to connect to Server!!"
                                                        Log.d(LOGTAG, "✅ Server connected!!")
                                                    }
                                                },
                                                onExit = {}
                                            )
                                        }
                                    },
                                    enabled = isRegistered && groupAddress == null,
                                ) {
                                    Text(text = "Scan QRCode")
                                }

                                Button(
                                    onClick = {
                                        coroutineScope.launch {
                                            // Step 5: Verify Identity Document
                                            Log.d(LOGTAG, "🔧 Open Identity document verification flow...")
                                            account?.openDocumentVerificationFlow(
                                                isDevMode = true,
                                                onFinish = { isSuccess, error ->
                                                    if (isSuccess) {
                                                        isDocumentVerified = true
                                                        statusText = "Your document is verified!!"
                                                        Log.d(LOGTAG, "✅ Identity Document is verified!!")
                                                    }
                                                }
                                            )
                                        }
                                    },
                                    enabled = isRegistered && groupAddress != null && !isDocumentVerified
                                ) {
                                    Text(text = "Verify Identity Document")
                                }

                                Button(
                                    onClick = {
                                        Log.d(LOGTAG, "🔧 Start sharing credentials...")
                                        Common.notifyServerForRequest( account ?: return@Button, groupAddress ?: return@Button, message = "PROVIDE_CREDENTIAL_DOCUMENT")
                                    },
                                    enabled = isRegistered && isDocumentVerified && groupAddress != null
                                ) {
                                    Text(text = "Start Sharing Document Credential")
                                }

                                Text(text = statusText)

                                // display credential request with buttons to confirm or reject.
                                if (requestMessage != null) {
                                    Dialog(
                                        onDismissRequest = { },
                                        properties = DialogProperties(dismissOnBackPress = false, dismissOnClickOutside = false, usePlatformDefaultWidth = false),
                                    ) {
                                        // Step 6: Share Identity Credential
                                        Log.d(LOGTAG, "✅ Display request message")
                                        account?.DisplayRequestUI(selfModifier, requestMessage ?: return@Dialog, onFinish = { isSent, status ->
                                            requestMessage = null
                                            Log.d(LOGTAG, "✅ Sharing identity credential successfully!")
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