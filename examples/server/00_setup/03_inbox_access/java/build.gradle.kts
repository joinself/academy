plugins {
    kotlin("jvm") version "2.0.21"
    application
}

repositories {
    mavenCentral()
    maven("https://central.sonatype.com/repository/maven-snapshots/")
}

dependencies {
    implementation("com.joinself:sdk-jvm:1.0.0")
}

kotlin {
    sourceSets.main {
        kotlin.srcDirs(".", "../../../common")
    }
}

application {
    mainClass = "MainKt"
} 
