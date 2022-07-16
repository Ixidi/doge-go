package objects

import (
	"doge/math"
	"doge/offsets"
	"doge/windows"
)

type Team byte

const (
	RedTeam     Team = 1
	BlueTeam    Team = 2
	NeutralTeam Team = 3
)

type GameObject struct {
	Index             uint32
	Team              Team
	Direction         math.Vector3
	Position          math.Vector3
	Dead              uint32
	Visibility        bool
	Mana              float32
	MaxMana           float32
	Invulnearable     bool
	Targetable        bool
	Health            float32
	MaxHealth         float32
	BonusAttackDamage float32
	AttackDamage      float32
	Armor             float32
	BonusArmor        float32
	MagicResist       float32
	MovementSpeed     float32
	AttackRange       float32
	ChampionName      string
}

func ReadGameObject(mem windows.Memory, address uint32) (obj GameObject, err error) {
	buff, err := mem.ReadBuff(12000, address)
	if err != nil {
		return
	}

	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()

	readFromBuffPanic(&obj.Index, offsets.ObjectIndex, buff)

	var team uint32
	readFromBuffPanic(&team, offsets.ObjectTeam, buff)
	switch team {
	case 100:
		obj.Team = BlueTeam
	case 200:
		obj.Team = RedTeam
	default:
		obj.Team = NeutralTeam
	}

	readVectorFromBuffPanic(&obj.Direction, offsets.ObjectDirection, buff)
	readVectorFromBuffPanic(&obj.Position, offsets.ObjectPosition, buff)
	readFromBuffPanic(&obj.Dead, offsets.ObjectDead, buff)
	readFromBuffPanic(&obj.Visibility, offsets.ObjectVisibility, buff)
	readFromBuffPanic(&obj.Mana, offsets.ObjectMana, buff)
	readFromBuffPanic(&obj.MaxMana, offsets.ObjectMaxMana, buff)
	readFromBuffPanic(&obj.Invulnearable, offsets.ObjectInvulnerable, buff)
	readFromBuffPanic(&obj.Targetable, offsets.ObjectTargetable, buff)
	readFromBuffPanic(&obj.Health, offsets.ObjectHealth, buff)
	readFromBuffPanic(&obj.MaxHealth, offsets.ObjectMaxHealth, buff)
	readFromBuffPanic(&obj.BonusAttackDamage, offsets.ObjectBonusAttackDamage, buff)
	readFromBuffPanic(&obj.AttackDamage, offsets.ObjectAttackDamage, buff)
	readFromBuffPanic(&obj.Armor, offsets.ObjectArmor, buff)
	readFromBuffPanic(&obj.BonusArmor, offsets.ObjectBonusArmor, buff)
	readFromBuffPanic(&obj.MagicResist, offsets.ObjectMagicResist, buff)
	readFromBuffPanic(&obj.MovementSpeed, offsets.ObjectMovementSpeed, buff)
	readFromBuffPanic(&obj.AttackRange, offsets.ObjectAttackRange, buff)
	readFromBuffPanic(&obj.ChampionName, offsets.ObjectChampionName, buff)
	return
}

func ReadGameObjectsInterfaceOffsets(mem windows.Memory, address uint32) ([]uint32, error) {
	var (
		arrayInterfaceAddress, arrayAddress, arraySize, championAddress uint32
	)
	if err := mem.Read(&arrayInterfaceAddress, address); err != nil {
		return nil, err
	}
	if err := mem.Read(&arrayAddress, arrayInterfaceAddress+0x04); err != nil {
		return nil, err
	}
	if err := mem.Read(&arraySize, arrayInterfaceAddress+0x08); err != nil {
		return nil, err
	}

	objects := make([]uint32, arraySize)
	for i := 0; i < int(arraySize); i++ {
		if err := mem.Read(&championAddress, arrayAddress+(uint32(i)*4)); err != nil {
			return nil, err
		}
		objects[i] = championAddress
	}

	return objects, nil
}
