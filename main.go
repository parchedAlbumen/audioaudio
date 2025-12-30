package main

import (
	"fmt"
	"math/cmplx"
	"os"

	"github.com/go-audio/wav"
	"github.com/mjibson/go-dsp/fft"
)

func main() {
	//get file here
	f, err := os.Open("audioFiles/test.wav")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	//decoder for the wav file
	decoder := wav.NewDecoder(f)
	if !decoder.IsValidFile() {
		panic("the wav file is bad")
	}

	//convert to PCM (the amplitudes basically)
	buf, err := decoder.FullPCMBuffer()

	//get some info here
	sampleRate := buf.Format.SampleRate
	channels := buf.Format.NumChannels

	floatVerBuf := buf.AsFloat32Buffer()
	samples := make([]float64, len(floatVerBuf.Data))
	for i, v := range floatVerBuf.Data {
		samples[i] = float64(v) //convert everything from float32 to float64
	}

	fmt.Println("sample rate: ", sampleRate)
	fmt.Println("# of channels: ", channels)
	fmt.Println("length: ", len(samples))

	//now we have sample rate, channels, length of sample, and the samples in float64
	//do fft next

	const frameSize = 2048

	frame := samples[1024 : frameSize+1024] //grabs the first frame
	spectrum := fft.FFTReal(frame)

	fmt.Println("\nFrequency spectrum:")
	for i := 0; i < len(spectrum)/2; i++ {
		mag := cmplx.Abs(spectrum[i])
		freq := float64(i) * float64(sampleRate) / float64(frameSize)
		fmt.Printf("Bin %3d (%7.1f Hz): %.2f\n", i, freq, mag)
	}
}
