import { HexColorPicker } from 'react-colorful'
import { useState } from 'react'
import {
  Popover,
  Button,
  PopoverTrigger,
  PopoverContent,
  PopoverArrow,
  PopoverBody,
  PopoverHeader,
  PopoverCloseButton,
  Box,
} from '@chakra-ui/react'

import 'react-colorful/dist/index.css'

const Picker = () => {
  const [color, setColor] = useState('#aabbcc')
  return (
    <Popover>
      <PopoverTrigger>
      <Box as="button" borderRadius="md" bg={color} color="white" px={4} h={8}>
        Select Background Color
        </Box>
      </PopoverTrigger>
      <PopoverContent>
        <PopoverArrow />
        <PopoverCloseButton />
        <PopoverHeader>Select Background Color!</PopoverHeader>
        <PopoverBody>
          <HexColorPicker color={color} onChange={setColor} />
        </PopoverBody>
      </PopoverContent>
    </Popover>
  )
}

export default Picker
