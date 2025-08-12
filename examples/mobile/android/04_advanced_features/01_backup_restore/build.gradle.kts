
plugins {
    alias(libs.plugins.android.application)
    alias(libs.plugins.kotlin.android)
    alias(libs.plugins.kotlin.compose)
}

android {
    namespace = "com.joinself.app.academy"
    compileSdk = libs.versions.compileSdk.get().toInt()

    defaultConfig {
        applicationId = "com.joinself.app.academy"
        minSdk = libs.versions.minSdk.get().toInt()
        targetSdk = libs.versions.targetSdk.get().toInt()
        versionCode = 1
        versionName = "1.0"

        ndk {
            abiFilters.clear()
            abiFilters.add("arm64-v8a")
        }
        setProperty("archivesBaseName", "self-academy")
    }

    compileOptions {
        sourceCompatibility = JavaVersion.VERSION_17
        targetCompatibility = JavaVersion.VERSION_17
    }
    kotlinOptions {
        jvmTarget = "17"
    }
    buildFeatures {
        compose = true
    }
    signingConfigs {
        getByName("debug") {
            storeFile = file("${rootDir.absolutePath}/debug.keystore")
        }
    }
    buildTypes {
        debug {
            isMinifyEnabled = false
            signingConfig = signingConfigs.getByName("debug")
        }
    }
    packaging {
        jniLibs {
            pickFirsts.addAll(listOf("lib/x86/libc++_shared.so", "lib/x86_64/libc++_shared.so", "lib/armeabi-v7a/libc++_shared.so", "lib/arm64-v8a/libc++_shared.so"))
            useLegacyPackaging = true
        }
        resources {
            excludes.addAll(listOf("META-INF/NOTICE", "META-INF/LICENSE", "META-INF/DEPENDENCIES", "META-INF/versions/9/OSGI-INF/MANIFEST.MF"))
        }
        dex {
            useLegacyPackaging = true
        }
    }
    dependenciesInfo {
        includeInApk = false
        includeInBundle = false
    }
}

dependencies {
    implementation(project(":common"))
    implementation(libs.self.sdk.android)

    implementation(libs.androidx.core.ktx)
    implementation(libs.androidx.lifecycle.runtime.ktx)
    implementation(libs.androidx.activity.compose)
    implementation(platform(libs.androidx.compose.bom))
    implementation(libs.androidx.ui)
    implementation(libs.androidx.ui.graphics)
    implementation(libs.androidx.ui.tooling.preview)
    implementation(libs.androidx.material3)
    implementation(libs.androidx.compose.navigation)
}
