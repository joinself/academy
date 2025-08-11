package com.joinself.app.academy

import Common
import android.os.Bundle
import android.util.Log
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.activity.enableEdgeToEdge
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.PaddingValues
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.items
import androidx.compose.material3.Button
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Scaffold
import androidx.compose.material3.Text
import androidx.compose.material3.TextField
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateListOf
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
import com.joinself.sdk.models.ChatMessage
import com.joinself.sdk.models.CredentialMessage
import com.joinself.sdk.models.CredentialRequest
import com.joinself.sdk.models.DataObject
import com.joinself.sdk.models.Message
import com.joinself.sdk.models.PublicKey
import com.joinself.sdk.models.Receipt
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
        val storagePath = File(applicationContext.filesDir.absolutePath + "/chat_basic")
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
                    val messages = remember { mutableStateListOf<String>() }
                    var inputMessage by remember { mutableStateOf("") }

                    fun sendChat() {
                        // build a attachment if there is one
                        val attachment = DataObject.Builder()
                            .setData("hello".encodeToByteArray())
                            .setMimeType("text/plain")
                            .build()

                        // build a chat message to send
                        val chat = ChatMessage.Builder()
                            .setMessage(inputMessage)
                            .setAttachments(listOf(attachment))
                            .build()

                        // send chat to server
                        coroutineScope.launch(Dispatchers.IO) {
                            account?.send(toAddress = groupAddress!!, message = chat)

                            messages.add(inputMessage)
                            inputMessage = ""
                        }
                    }

                    LaunchedEffect(true) {
                        account = Account.Builder()
                            .setContext(applicationContext)
                            .setEnvironment(Environment.production)
                            .setSandbox(true)
                            .setStoragePath(storagePath.absolutePath)
                            .setCallbacks(object : Account.Callbacks {
                                override fun onMessage(message: Message) {
                                    Log.d(LOGTAG, "onMessage: ${message.id()}")
                                    if (message is ChatMessage) {
                                        messages.add(message.message()) // append text to the message list
                                    }
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
                                    enabled = isRegistered && groupAddress == null,
                                ) {
                                    Text(text = "Scan QRCode")
                                }

                                Text(text = statusText)

                                Row(
                                    horizontalArrangement = Arrangement.spacedBy(4.dp),
                                    verticalAlignment = Alignment.CenterVertically
                                ) {
                                    TextField(modifier = Modifier.weight(1f),
                                        value = inputMessage,
                                        onValueChange = {
                                            inputMessage = it
                                        },
                                        enabled = groupAddress != null,
                                        placeholder = { Text("input a message") }
                                    )
                                    Button(modifier = Modifier.width(80.dp), contentPadding = PaddingValues(0.dp),
                                        onClick = {
                                            sendChat()
                                        },
                                        enabled = isRegistered && groupAddress != null
                                    ) {
                                        Text(text = "Send")
                                    }
                                }
                                LazyColumn(modifier = Modifier.fillMaxSize().background(Color.LightGray)) {
                                    items(messages) { msg ->
                                        Text(
                                            text = msg,
                                            modifier = Modifier.padding(start = 4.dp)
                                        )
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