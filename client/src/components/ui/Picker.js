import { HexColorPicker } from 'react-colorful'
import { useState, useEffect } from 'react'
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



const Picker = (props) => {
  const [color, setColor] = useState('#aabbcc')
  const handleChange = (color) => {
    setColor(color)
    props.setValue(props.name, color);
  } 
  useEffect(() => {
    props.register(props.name); // custom register Antd input
    props.setValue(props.name, color);
  }, [props.register])  
  return (
    <Popover>
      <PopoverTrigger>
      <Box as="button" type="button" borderRadius="md" bg={color} color="white" px={4} h={8}>
        {props.text}
        </Box>
      </PopoverTrigger>
      <PopoverContent>
        <PopoverArrow />
        <PopoverCloseButton />
        <PopoverHeader>{props.text}!</PopoverHeader>
        <PopoverBody>
          <HexColorPicker color={color} onChange={handleChange} />
        </PopoverBody>
      </PopoverContent>
    </Popover>
  )
}

export default Picker
