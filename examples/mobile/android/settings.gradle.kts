pluginManagement {
    repositories {
        google {
            content {
                includeGroupByRegex("com\\.android.*")
                includeGroupByRegex("com\\.google.*")
                includeGroupByRegex("androidx.*")
            }
        }
        mavenCentral()
        gradlePluginPortal()
    }
}
dependencyResolutionManagement {
    repositories {
        maven { url = uri("https://central.sonatype.com/repository/maven-snapshots/") }
        google()
        mavenCentral()
    }
}

rootProject.name = "SelfAcademy"
include(":00_setup:01_new_account")