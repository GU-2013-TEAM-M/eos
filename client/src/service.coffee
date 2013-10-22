# Sets a cookie
# Params:	name - cookie name
#			value - cookie value
#			days - days until a cookie should expire
setCookie = (name, value, days) ->
	if days
		date = new Date
		date.setTime(date.getTime() + (days*24*60*60*1000))
		expires = "; expires=" + date.toGMTString()
	else
		expires = ""
	document.cookie = name + "=" + value + expires + "; path=/"

# Gets a cookie value
# Params:	key - cookie name
# Return:	cookie value or null if the cookie was not found
getCookie = (key) ->
	key = key + "="
	for c in document.cookie.split(';')
		c.substring(1, c.length) while c.charAt(0) is ' '
		return c.substring(key.length, c.length) if c.indexOf(key) == 0
	return null