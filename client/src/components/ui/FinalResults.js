import React from 'react'
import { Box, Heading, Stack, Button } from '@chakra-ui/react'
import { ViewIcon, DownloadIcon } from '@chakra-ui/icons'

export default function FinalResults ({ response }) {
  return (
    <>
      <Box textAlign='center' pb={6}>
        <Heading>Finished! Download Now </Heading>
      </Box>
      <Box boxSize='sm'>
        <Stack direction='row' spacing={4}>
          <Button leftIcon={<ViewIcon />} onClick={()=> window.open(response.ResultURLs[0], "_blank")} variant='outline'>
            Preview
          </Button>
          <Button rightIcon={<DownloadIcon />} variant='outline'>
            Download
          </Button>
        </Stack>
      </Box>
    </>
  )
}
