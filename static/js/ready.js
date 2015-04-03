var map = {13: false, 16: false}

$(document)
	.ready(function() {
		console.log("loaded ready.js...")
		linesCode()
		$("textarea").allowTabChar()
		$("textarea").keyup(linesCode)
		$(".output").height($(document).height())
		
		$("a#run").click(function(e) {
			e.preventDefault()
	        $("form#script").submit()
		})

		$("a#save").click(function(e) {
			e.preventDefault()
			var filepath = prompt('Save as...')
			$.post("/save", { textarea: $("textarea").val(), filepath: filepath }, function(result) {
				console.log(result)
			})
		})

	})
	.keydown(function(e) {
	    if (e.keyCode in map) {
	        map[e.keyCode] = true;
	        if (map[13] && map[16]) {
				e.preventDefault()
				console.log("submitting script")
	            // $("button#submit").click()
	            $("form#script").submit()

	        }
	    }
	}).keyup(function(e) {
	    if (e.keyCode in map) {
	        map[e.keyCode] = false;
	    }
	})

function linesCode () {
	var linenums = $("div.linenums")
	linenums.html("")

	var frame = $("iframe"),
	contents = frame.contents(),
	body = contents.find("body")
	// var lines = body.html().match(/<\/div>/gmi)

	// var code = $("iframe#editor").html()
	// var lines = code.match(/\n/gmi)
	
	// for (var i = 1; i <= lines.length+50; i++) {
	for (var i = 1; i <= 60; i++) {

		linenums.append("<div class=\"num\">"+i+"</div>")
	}
}
