package app.iptocidr

import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.material.*
import androidx.compose.runtime.*
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp

@Composable
fun App() {
    var ip by remember { mutableStateOf("") }

    fun convert() {
        println(ip)
    }

    MaterialTheme {
        Column(Modifier.fillMaxWidth(), horizontalAlignment = Alignment.CenterHorizontally) {
            Spacer(modifier = Modifier.height(60.dp))
            Text("Type in an IP address", fontSize = 40.sp, fontWeight = FontWeight.ExtraBold)
            Spacer(modifier = Modifier.height(20.dp))
            OutlinedTextField(
                value = ip,
                onValueChange = { ip = it },
                placeholder = { Text("000.00.000.00") },
            )
            Spacer(modifier = Modifier.height(20.dp))
            Button(onClick = { convert() }) {
                Text("Convert")
            }
            Spacer(modifier = Modifier.height(60.dp))
        }
    }
}