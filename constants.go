package gobark

type LevelType string

type SoundType string

//goland:noinspection SpellCheckingInspection
const (
	DefaultDomain      string    = "api.day.app"
	DefaultSound       SoundType = Alarm
	Alarm              SoundType = "alarm"
	Anticipate         SoundType = "anticipate"
	Bell               SoundType = "bell"
	Birdsong           SoundType = "birdsong"
	Bloom              SoundType = "bloom"
	Calypso            SoundType = "calypso"
	Chime              SoundType = "chime"
	Choo               SoundType = "choo"
	Descent            SoundType = "descent"
	Electronic         SoundType = "electronic"
	Fanfare            SoundType = "fanfare"
	Glass              SoundType = "glass"
	GotoSleep          SoundType = "gotosleep"
	HealthNotification SoundType = "healthnotification"
	Horn               SoundType = "horn"
	Ladder             SoundType = "ladder"
	MailSend           SoundType = "mailsend"
	Minuet             SoundType = "minuet"
	MultiwayInvitation SoundType = "multiwayinvitation"
	NewMail            SoundType = "newmail"
	NewsFlash          SoundType = "newsflash"
	Noir               SoundType = "noir"
	PaymentSuccess     SoundType = "paymentsuccess"
	Shake              SoundType = "shake"
	SherwoodForest     SoundType = "sherwoodforest"
	Spell              SoundType = "spell"
	Suspense           SoundType = "suspense"
	Telegraph          SoundType = "telegraph"
	Tiptoes            SoundType = "tiptoes"
	Typewriters        SoundType = "typewriters"
	Update             SoundType = "update"
	Active             LevelType = "active"
	TimeSensitive      LevelType = "timeSensitive"
	Passive            LevelType = "passive"
)
