var map = {13: false, 16: false}

$(document)
	.ready(function() {
		console.log("loaded ready.js")
		linesCode()
		$("textarea").allowTabChar()
		$("textarea").keyup(linesCode)
			$(".output").height($(document).height())				
	})
	.keydown(function(e) {
	    if (e.keyCode in map) {
	        map[e.keyCode] = true;
	        if (map[13] && map[16]) {
				e.preventDefault()
	            $("button#editor").click()
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
	var code = $("textarea").val()
	var lines = code.match(/\n/gmi)
	for (var i = 1; i <= lines.length+50; i++) {
		linenums.append("<div class=\"num\">"+i+"</div>")
	}
}
