package objects

import (
	"doge/math"
	"doge/offsets"
	"doge/win"
)

type GameObject struct {
	Index             uint32
	Team              int32
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

func ReadGameObject(mem win.Memory, address uint32) (obj GameObject, err error) {
	buff, err := mem.ReadBuff(12000, address)
	if err != nil {
		return
	}

	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()

	readPanic(&obj.Index, offsets.ObjectIndex, buff)
	readPanic(&obj.Team, offsets.ObjectTeam, buff)
	readVectorPanic(&obj.Direction, offsets.ObjectDirection, buff)
	readVectorPanic(&obj.Position, offsets.ObjectPosition, buff)
	readPanic(&obj.Dead, offsets.ObjectDead, buff)
	readPanic(&obj.Visibility, offsets.ObjectVisibility, buff)
	readPanic(&obj.Mana, offsets.ObjectMana, buff)
	readPanic(&obj.MaxMana, offsets.ObjectMaxMana, buff)
	readPanic(&obj.Invulnearable, offsets.ObjectInvulnerable, buff)
	readPanic(&obj.Targetable, offsets.ObjectTargetable, buff)
	readPanic(&obj.Health, offsets.ObjectHealth, buff)
	readPanic(&obj.MaxHealth, offsets.ObjectMaxHealth, buff)
	readPanic(&obj.BonusAttackDamage, offsets.ObjectBonusAttackDamage, buff)
	readPanic(&obj.AttackDamage, offsets.ObjectAttackDamage, buff)
	readPanic(&obj.Armor, offsets.ObjectArmor, buff)
	readPanic(&obj.BonusArmor, offsets.ObjectBonusArmor, buff)
	readPanic(&obj.MagicResist, offsets.ObjectMagicResist, buff)
	readPanic(&obj.MovementSpeed, offsets.ObjectMovementSpeed, buff)
	readPanic(&obj.AttackRange, offsets.ObjectAttackRange, buff)
	readPanic(&obj.ChampionName, offsets.ObjectChampionName, buff)

	return
}
