package offsets

/*
Dump done: 2022-07-14 01:17:37
Live patch: 12.13.4533037+branch.releases-12-13.code.public
*/

const ( // Basic Offsets
	LocalPlayer      = 0x310ed68
	GameTime         = 0x31081a8
	ObjectManager    = 0x1870bd8
	Renderer         = 0x3141584
	UnderMouseObject = 0x24bfc24
	ViewProjMatrices = 0x313e6a0 // +0 view +64 projection
)

const ( // Object related Offsets
	ObjectIndex             = 0x0008 // Dumper could not find this offset.
	ObjectTeam              = 0x0034 // Dumper could not find this offset.
	ObjectDirection         = 0x0080 // Dumper could not find this offset.
	ObjectPosition          = 0x01DC // Dumper could not find this offset.
	ObjectDead              = 0x021C // Dumper could not find this offset.
	ObjectVisibility        = 0x0274 // Dumper could not find this offset.
	ObjectMana              = 0x029C // Dumper could not find this offset.
	ObjectMaxMana           = 0x02AC // Dumper could not find this offset.
	ObjectInvulnerable      = 0x03D4 // Dumper could not find this offset.
	ObjectTargetable        = 0x0D04 // Dumper could not find this offset.
	ObjectHealth            = 0x0E74 // Dumper could not find this offset.
	ObjectMaxHealth         = 0x0E84 // Dumper could not find this offset.
	ObjectBonusAttackDamage = 0x12C4 // Dumper could not find this offset.
	ObjectAttackDamage      = 0x134C // Dumper could not find this offset.
	ObjectArmor             = 0x1374 // Dumper could not find this offset.
	ObjectBonusArmor        = 0x1378 // Dumper could not find this offset.
	ObjectMagicResist       = 0x137C // Dumper could not find this offset.
	ObjectMovementSpeed     = 0x138C // Dumper could not find this offset.
	ObjectAttackRange       = 0x1394 // Dumper could not find this offset.
	ObjectChampionName      = 0x2BA4 // Dumper could not find this offset.
)

const ( // SpellBook related Offsets
)

const ( // Buffs related Offsets
)

const ( // Object lists
	HeroInterface = 0x1870c68
)

const ( // Renderer related Offsets
)

const ( // Utility Offsets
	GameVersion = 0x1510f04
)
