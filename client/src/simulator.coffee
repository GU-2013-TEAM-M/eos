simulate = (timing) ->
	new Graph("simulation", "karma");
	new Graph("simulation", "cpu", {count: 4});
	new Graph("simulation", "ram", {total: 2048});

	karmaGenerator = new SmoothRandomGenerator(100, 10)
	cpuGenerator1 = new SmoothRandomGenerator(100, 10)
	cpuGenerator2 = new SmoothRandomGenerator(100, 10)
	cpuGenerator3 = new SmoothRandomGenerator(100, 10)
	cpuGenerator4 = new SmoothRandomGenerator(100, 10)
	ramGenerator = new SmoothRandomGenerator(2048, 10)

	setInterval () ->
		karmaMessage = 
			type: "monitoring"
			data: {"daemon_id": "simulation", "data": {"karma": karmaGenerator.getNumber(), "cpu": [cpuGenerator1.getNumber(), cpuGenerator2.getNumber(), cpuGenerator3.getNumber(), cpuGenerator4.getNumber()], "ram": ramGenerator.getNumber()}}
		trySendMessage karmaMessage
	, timing


class SmoothRandomGenerator
	constructor: (max, diff) ->
		@max = max
		@current = max / 2
		@diff = diff

	getNumber: () ->
		@current += (if Math.random() > 0.5 then 1 else -1) * (Math.random()*@max+1)/@diff
		if @current > @max
			@current = @max
		if @current < 0
			@current = 0
		return Math.floor @current

	getNumbers: (n) ->
		numbers = []
		for i in [0...n]
			numbers.push(@getNumber())
		return numbers