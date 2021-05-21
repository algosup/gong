package main

import (
	"image/color"

	game "github.com/algosup/gamev2"
	"github.com/algosup/gamev2/key"
)

const width = 800
const height = 600

const padWidth = 15
const padHeight = 80
const padSpeed = 5

const dotSize = 15
const ballSize = 15

const leftPadX = padWidth * 2
const rightPadX = width - padWidth*3

const scoreSize = 10

var leftPadY = (height - padHeight) / 2
var rightPadY = (height - padHeight) / 2

var ballSpeedX = 2
var ballSpeedY = 2

var ballX = width / 2
var ballY = height / 2

var leftScore = 0
var rightScore = 0

var numbers = [][][]byte{
	{
		{1, 1, 1},
		{1, 0, 1},
		{1, 0, 1},
		{1, 0, 1},
		{1, 1, 1},
	},
	{
		{0, 1, 0},
		{1, 1, 0},
		{0, 1, 0},
		{0, 1, 0},
		{1, 1, 1},
	},
	{
		{1, 1, 1},
		{0, 0, 1},
		{1, 1, 1},
		{1, 0, 0},
		{1, 1, 1},
	},
	{
		{1, 1, 1},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 1},
		{1, 1, 1},
	},
	{
		{1, 0, 1},
		{1, 0, 1},
		{1, 1, 1},
		{0, 0, 1},
		{0, 0, 1},
	},
	{
		{1, 1, 1},
		{1, 0, 0},
		{1, 1, 1},
		{0, 0, 1},
		{1, 1, 1},
	},
	{
		{1, 1, 1},
		{1, 0, 0},
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	},
	{
		{1, 1, 1},
		{0, 0, 1},
		{0, 0, 1},
		{0, 0, 1},
		{0, 0, 1},
	},
	{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	},
	{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
		{0, 0, 1},
		{1, 1, 1},
	},
}

func drawTerrain() {
	var y = 0
	for y < height {
		game.DrawRect((width-dotSize)/2, y, dotSize, dotSize, color.White)
		y = y + dotSize*2
	}
}

func checkKeyboard() {
	if key.IsPressed(key.Shift) {
		if leftPadY > 0 {
			leftPadY -= padSpeed
		}
	}

	if key.IsPressed(key.Control) {
		if leftPadY < height-padHeight {
			leftPadY += padSpeed
		}
	}

	if key.IsPressed(key.Up) {
		if rightPadY > 0 {
			rightPadY -= padSpeed
		}
	}

	if key.IsPressed(key.Down) {
		if rightPadY < height-padHeight {
			rightPadY += padSpeed
		}
	}
}

func drawBall() {
	game.DrawRect(ballX, ballY, ballSize, ballSize, color.White)
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
func intersect(start1, end1, start2, end2 int) bool {
	start := max(start1, start2)
	end := min(end1, end2)
	return end > start
}

func intersectRect(
	left1,
	top1,
	right1,
	bottom1,
	left2,
	top2,
	right2,
	bottom2 int) bool {
	return intersect(left1, right1, left2, right2) &&
		intersect(top1, bottom1, top2, bottom2)
}

func moveBall() {
	ballX += ballSpeedX
	ballY += ballSpeedY

	if ballY <= 0 {
		ballSpeedY = -ballSpeedY
	}

	if ballY >= height-ballSize {
		ballSpeedY = -ballSpeedY
	}

	if intersectRect(ballX, ballY, ballX+ballSize, ballY+ballSize,
		leftPadX, leftPadY, leftPadX+padWidth, leftPadY+padHeight) {
		ballSpeedX = -ballSpeedX
	}

	if intersectRect(ballX, ballY, ballX+ballSize, ballY+ballSize,
		rightPadX, rightPadY, rightPadX+padWidth, rightPadY+padHeight) {
		ballSpeedX = -ballSpeedX
	}

	if ballX < -ballSize {
		rightScore++
		ballX = width / 2
		ballY = height / 2
	}

	if ballX > width {
		leftScore++
		ballX = width / 2
		ballY = height / 2
	}

}

func drawScore(score int, left int, top int) {
	if left > width/2 {
		s := score

		for s >= 10 {
			left += scoreSize * 4
			s /= 10
		}
	}

	n := numbers[score%10]

	for j, r := range n {
		for i, c := range r {
			if c == 1 {
				game.DrawRect(left+i*scoreSize, top+j*scoreSize, scoreSize, scoreSize, color.White)
			}
		}
	}

	if score >= 10 {
		drawScore(score/10, left-scoreSize*4, top)
	}
}

func draw() {
	drawTerrain()
	drawBall()
	drawScore(leftScore, width/2-scoreSize*5, scoreSize)
	drawScore(rightScore, width/2+scoreSize*2, scoreSize)
	game.DrawRect(leftPadX, leftPadY, padWidth, padHeight, color.White)
	game.DrawRect(rightPadX, rightPadY, padWidth, padHeight, color.White)

	moveBall()
	checkKeyboard()
}

func main() {
	game.Run("Gong", width, height, draw)
}
