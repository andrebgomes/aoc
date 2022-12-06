package twok

func Six(input string) ([2]interface{}, error) {
	startOfPacket := 0
	startOfMessage := 0

	packetStream := []string{}
	messageStream := []string{}
	for k, char := range input {
		charStr := string(char)

		if k > 3 {
			if startOfPacket == 0 {
				if !isThereTwo(packetStream) {
					startOfPacket = k
				}

				lastThree := packetStream[1:]
				packetStream = append(lastThree, charStr)
			}
		} else {
			packetStream = append(packetStream, charStr)
			messageStream = append(messageStream, charStr)
			continue
		}

		if k > 13 {
			if startOfMessage == 0 {
				if !isThereTwo(messageStream) {
					startOfMessage = k
				}

				lastThirteen := messageStream[1:]
				messageStream = append(lastThirteen, charStr)
			}
		} else {
			messageStream = append(messageStream, charStr)
			continue
		}

		if startOfPacket != 0 && startOfMessage != 0 {
			break
		}
	}

	return [2]interface{}{startOfPacket, startOfMessage}, nil
}

func isThereTwo(array []string) bool {
	existing := map[string]bool{}
	for _, r := range array {
		if existing[r] {
			return true
		}
		existing[r] = true
	}
	return false
}
