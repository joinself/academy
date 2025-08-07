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
include(":common")
include(":00_setup:01_new_account")
include(":01_connection:01_address")
include(":01_connection:02_qr")
include(":02_credentials:01_get_credentials:01_verify_email")
include(":02_credentials:01_get_credentials:02_verify_identity_document")
include(":02_credentials:01_get_credentials:03_get_custom_credential")
include(":02_credentials:02_share_credentials:01_authentication")
include(":02_credentials:02_share_credentials:02_email_credential")
include(":02_credentials:02_share_credentials:03_document_credential")
include(":02_credentials:03_digitial_signatures")