package hunt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSharkHuntsSuccessfully(t *testing.T) {

	//arange
	shark := Shark{
		hungry: true,
		tired:  false,
		speed:  10,
	}

	prey := Prey{
		name:  "pececito",
		speed: 5,
	}

	sharkExpected := Shark{
		hungry: false,
		tired:  true,
		speed:  10,
	}

	//act
	resultHunt := shark.Hunt(&prey)

	//assert

	assert.Empty(t, resultHunt)
	assert.Equal(t, sharkExpected, shark)
}

func TestSharkCannotHuntBecauseIsTired(t *testing.T) {
	//arange
	shark := Shark{
		hungry: true,
		tired:  true,
		speed:  10,
	}

	prey := Prey{
		name:  "pececito",
		speed: 5,
	}

	//act
	resultHunt := shark.Hunt(&prey)

	//assert

	assert.Error(t, resultHunt)
	assert.EqualError(t, resultHunt, "cannot hunt, i am really tired")

}

func TestSharkCannotHuntBecaisIsNotHungry(t *testing.T) {
	//arange
	shark := Shark{
		hungry: false,
		tired:  false,
		speed:  10,
	}

	prey := Prey{
		name:  "pececito",
		speed: 5,
	}

	//act
	resultHunt := shark.Hunt(&prey)

	//assert

	assert.Error(t, resultHunt)
	assert.EqualError(t, resultHunt, "cannot hunt, i am not hungry")
}

func TestSharkCannotReachThePrey(t *testing.T) {
	//arange
	shark := Shark{
		hungry: true,
		tired:  false,
		speed:  5,
	}

	prey := Prey{
		name:  "pececito",
		speed: 10,
	}

	//act
	resultHunt := shark.Hunt(&prey)

	//assert

	assert.Error(t, resultHunt)
	assert.EqualError(t, resultHunt, "could not catch it")

}

func TestSharkHuntNilPrey(t *testing.T) {
	//arange
	shark := Shark{
		hungry: true,
		tired:  false,
		speed:  5,
	}

	var prey *Prey = nil

	//act
	resultHunt := shark.Hunt(prey)

	//assert

	assert.Error(t, resultHunt)
	assert.EqualError(t, resultHunt, "shark Hunt Nil Prey")
}
