import React from 'react'
import { Box, Image } from '@chakra-ui/react'

export default function ScreenshotPreview ({url}) {
  return (
    <Box boxSize='sm'>
      <Image src={url} alt='Segun Adebayo' />

    </Box>
  )
}
