package process

const sar_dir = "/storage/emulated/0/SAR/"

var (
	ignoring_path = []string{
		"PacProcessor",
		"SystemUITools",
	}
)

/*
////app

020a0000000000000000000000000000.drbin
AlarmClock
AppCenter
AricentIMSService
Bluetooth
BluetoothMidiService
Calculator
Calendar
Camera
CaptivePortalLogin
CertInstaller
ConnectivitySettings
DMSettings
DesktopBackup
DocumentsUI
DownloadProviderUi
EngineerMode
FileManager
FlymeCommunication
FlymeTelecomm
GFManager
HTMLViewer
ImsATCommandTester
LRA_Soft
LRA_Strong
LockScreenFramework
Map
MeizuPay
Morelocale
MzAccount
MzBatteryManager
MzCloudService
MzPhoneLocationService
MzSecurity
MzSetupWizard
MzShare
MzSimContacts
MzSyncService
MzUpdate
NetAdmin
NfcNci
NfceeService
NotePaper
PacProcessor
PaymentSwitch
PrivacyController
Search
SecJniRil
SmartcardService
SplitLocController
Stk
SysDebugMode
SystemWallpaper
ToolBox
UARTSwitch
UPTsmService
USBModeSwitch
UserDictionaryProvider
WAPPushManager
WalletService
Weather
com.wolfsonmicro.ez2control
com.wolfsonmicro.ez2control.dummyrecognizer
mcRegistry
vlife-sdk-mx
webview

////priv-app

AMapNetworkLocation_meizu
ActivatePhone
AlphaMe
BackupRestoreConfirmation
BlackList
CalendarProvider
CallLogBackup
CarrierConfig
ChildrenLauncher
ContactsProvider
CustomizeCenter
DefaultContainerService
Dialer
DownloadProvider
EasyLauncher
Email
ExperienceDataSync
ExternalStorageProvider
FlymeLauncher
FusedLocation
Gallery
HomeTool
InCallUI
InputDevices
LauncherMenuProvider
ManagedProvisioning
MediaProvider
Mms
MmsService
MzAccountPlugin
MzBackup
MzCallSetting
MzInput
NetContactService
OneTimeInitializer
PackageInstaller
PerfUI
PowerSaveModeHome
ProxyHandler
RemoteCooperation
SceneInfo
Settings
SettingsProvider
SharedStorageBackup
Shell
SoundRecorder
StatementService
SystemUI
SystemUITools
Tag
TeleService
Telecom
TelephonyProvider
VpnDialogs
WallpaperCropper
 */