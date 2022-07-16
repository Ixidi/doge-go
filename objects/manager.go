package objects

import (
	"doge/offsets"
	"doge/windows"
)

type ObjectManager interface {
	LocalPlayer() *GameObject
	Champions() []*GameObject
	Game() *Game
	Reread() error
}

type objectManager struct {
	mem                windows.Memory
	localPlayerAddress uint32
	championAddresses  []uint32
	localPlayer        *GameObject
	champions          []*GameObject
	game               *Game
}

func (o objectManager) LocalPlayer() *GameObject {
	return o.localPlayer
}

func (o objectManager) Champions() []*GameObject {
	return o.champions
}

func (o objectManager) Game() *Game {
	return o.game
}

func (o *objectManager) Reread() error {
	localPlayer, err := ReadGameObject(o.mem, o.localPlayerAddress)
	if err != nil {
		return err
	}

	for i, address := range o.championAddresses {
		object, err := ReadGameObject(o.mem, address)
		if err != nil {
			return err
		}
		o.champions[i] = &object
	}

	o.localPlayer = &localPlayer
	game, err := ReadGame(o.mem)
	if err != nil {
		return err
	}
	o.game = &game
	return nil
}

func NewObjectManager(mem windows.Memory) (ObjectManager, error) {
	baseAddress := uint32(mem.Process().BaseAddress)

	var objectManager objectManager
	objectManager.mem = mem
	if err := mem.Read(&objectManager.localPlayerAddress, baseAddress+offsets.LocalPlayer); err != nil {
		return nil, err
	}

	championAddresses, err := ReadGameObjectsInterfaceOffsets(mem, baseAddress+offsets.HeroInterface)
	if err != nil {
		return nil, err
	}

	objectManager.championAddresses = championAddresses
	objectManager.champions = make([]*GameObject, len(championAddresses))
	err = objectManager.Reread()
	if err != nil {
		return nil, err
	}
	return &objectManager, nil
}
