/*

	This random number generator that we have right here is what we refer to as a pseudo random generator

	Go by default uses a random number generator that depends upon some seed value
	You can imagine the seed value as being kind of like the source of randomness inside of our number generator
	We take the seed value we pass in the generator and then the generator is used to make a random sequence of numbers or values or whatever we are in randomising
	So the issue here is that by default the Go random number generator always uses the exact same seed

	So every single time we restart as soon as we start trying to generate random numbers we are always going to get the exact same sequence because we always start from the same exact seed value

	So in order to fix this we need to somehow change this seed value
	We need to say every single timr that we start up our program; We generate some new seed value and use that as the seed to generate any random numbers in the future


	func (d deck) shuffle() {

		// Generate a different int64 every single time
		// We use that as the seed to generate a newSource object
		// And then we use that source object as the basis for our new random number generator
		source := rand.NewSource(time.Now().UnixNano())
		r := rand.New(source) //

		for index := range d {
			newPosition := r.Intn(len(d) - 1)

			d[index], d[newPosition] = d[newPosition], d[index]
		}
	}


*/