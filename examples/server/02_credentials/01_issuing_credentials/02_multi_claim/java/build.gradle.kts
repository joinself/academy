plugins {
    kotlin("jvm") version "2.0.21"
    application
}

repositories {
    mavenCentral()
    maven("https://central.sonatype.com/repository/maven-snapshots/")
}

dependencies {
    implementation("com.joinself:sdk-jvm:1.0.1")
}

kotlin {
    sourceSets.main {
        kotlin.srcDirs(".", "../../../../common")
    }
    jvmToolchain(17)
}

application {
    mainClass = "MainKt"
}

tasks.named<JavaExec>("run") {
    standardInput = System.`in`
}