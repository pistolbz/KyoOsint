package controllers

import (
	"goosint/modules/fbinfo"
	"goosint/modules/github"
	"goosint/modules/ipinfo"
	"goosint/modules/seon"
	"goosint/modules/whois"
	"net"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

func IPController(c *gin.Context) {
	input, check := c.GetPostForm("input")

	// validate input
	inputChecker(input, check, c)
	if !isValidIP(input) {
		c.JSON(http.StatusOK, gin.H{"msg": "You think this is IP?, you are wrong :))"})
		return
	}

	result := gin.H{}
	ipInfoResult := ipinfo.IPInfo(input)

	result["ipinfo"] = ipInfoResult

	c.JSON(http.StatusOK, result)
}

// phone
func PhoneController(c *gin.Context) {
	input, check := c.GetPostForm("input")

	// validate input
	inputChecker(input, check, c)
	if !isValidPhone(input) {
		c.JSON(http.StatusOK, gin.H{"msg": "You think this is phone?, you are wrong :))"})
		return
	}

	if strings.HasPrefix(input, "0") && len(input) == 10 {
		input = "84" + input[1:]
	}

	result := gin.H{}
	seonResult := seon.SeonPhoneSearch(input)
	domainRegisterByPhone := whois.ReverseWhois(input)

	if seonResult.Success {
		result["seon"] = seonResult
	}
	result["whois"] = domainRegisterByPhone

	c.JSON(http.StatusOK, result)
}

// email
func EmailController(c *gin.Context) {
	input, check := c.GetPostForm("input")

	// validate input
	inputChecker(input, check, c)
	if !isValidEmail(input) {
		c.JSON(http.StatusOK, gin.H{"msg": "You think this is email?, you are wrong :))"})
		return
	}

	result := gin.H{}
	usernameGithubList := github.GithubSearchEmail(input)
	seonResult := seon.SeonEmailSearch(input)
	domainRegisterByEmail := whois.ReverseWhois(input)

	result["github"] = usernameGithubList
	if seonResult.Success {
		result["seon"] = seonResult
	}
	result["whois"] = domainRegisterByEmail

	c.JSON(http.StatusOK, result)
}

// Username
func UsernameController(c *gin.Context) {
	input, check := c.GetPostForm("input")

	// validate input
	inputChecker(input, check, c)
	if !isValidUsername(input) {
		c.JSON(http.StatusOK, gin.H{"msg": "You think this is username?, you are wrong :))"})
		return
	}

	result := gin.H{}
	emailGithubList := github.GithubSearchName(input)

	result["github"] = emailGithubList

	c.JSON(http.StatusOK, result)
}

// facebook uid
func FacebookController(c *gin.Context) {
	input, check := c.GetPostForm("input")

	// validate input
	inputChecker(input, check, c)
	if !isValidFacebookUid(input) {
		c.JSON(http.StatusOK, gin.H{"msg": "You think this is Facebook UID?, you are wrong :))"})
		return
	}

	result := gin.H{}
	fbinfo := fbinfo.GetFacebookInfo(input)

	result["facebook"] = fbinfo

	c.JSON(http.StatusOK, result)
}

// handle controller
func inputChecker(input string, check bool, c *gin.Context) {
	if !check {
		c.JSON(http.StatusNotFound, gin.H{"msg": "Where's your input?"})
		return
	}
}

func isValidEmail(email string) bool {
	// Kiểm tra định dạng email bằng regular expression
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

func isValidUsername(username string) bool {
	// Kiểm tra định dạng tên người dùng bằng regular expression
	usernameRegex := regexp.MustCompile(`^[a-zA-Z0-9_-]{3,16}$`)
	return usernameRegex.MatchString(username)
}

func isValidPhone(phone string) bool {
	// Kiểm tra định dạng số điện thoại bằng regular expression
	phoneRegex := regexp.MustCompile(`^(0\d{9}|84\d{9})$`)
	return phoneRegex.MatchString(phone)
}

func isValidFacebookUid(uid string) bool {
	// Kiểm tra định dạng Facebook UID bằng regular expression
	uidRegex := regexp.MustCompile(`^\d{6,15}$`)
	return uidRegex.MatchString(uid)
}

func isValidIP(ip string) bool {
	return net.ParseIP(ip) != nil
}
