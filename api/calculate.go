package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	calculator "go-taylor/calculator"
)

type errorResponse struct {
	Error string `json:"error"`
}

type response struct {
	Result string `json:"result"`
}

func writeError(w http.ResponseWriter, err string) {
	res := errorResponse{err}

	json, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)
	w.Write(json)
}

func writeResponse(w http.ResponseWriter, result float64) {
	res := response{strconv.FormatFloat(result, 'f', -1, 64)}

	json, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "s-maxage=120, stale-while-revalidate")
	w.WriteHeader(200)
	w.Write(json)
}

func Handler(w http.ResponseWriter, r *http.Request) {

	commands, okCommand := r.URL.Query()["command"]
	inputs, okInput := r.URL.Query()["input"]

	if !okCommand || !okInput {
		writeError(w, "No command or no input")

		return
	}

	switch command := commands[0]; command {
	case "exp":
		{
			input, _ := strconv.ParseFloat(inputs[0], 64)
			result := calculator.Exp(input)
			writeResponse(w, result)
			break
		}
	case "norm":
		{
			input, _ := strconv.ParseFloat(inputs[0], 64)
			result := calculator.Norm(input)
			writeResponse(w, result)
			break
		}
	case "sin":
		{
			input, _ := strconv.ParseFloat(inputs[0], 64)
			result := calculator.Sin(input)
			writeResponse(w, result)
			break
		}
	case "cos":
		{
			input, _ := strconv.ParseFloat(inputs[0], 64)
			result := calculator.Cos(input)
			writeResponse(w, result)
			break
		}
	case "tan":
		{
			input, _ := strconv.ParseFloat(inputs[0], 64)

			result := calculator.Sin(input) / calculator.Cos(input)
			writeResponse(w, result)
			break
		}
	case "cot":
		{
			input, _ := strconv.ParseFloat(inputs[0], 64)

			result := calculator.Cos(input) / calculator.Sin(input)
			writeResponse(w, result)
			break
		}
	default:
		{
			writeError(w, "Invalid command")
		}
	}
}
