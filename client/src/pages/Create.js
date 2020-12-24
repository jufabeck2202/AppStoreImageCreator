import { Flex, Heading, Stack, Box } from '@chakra-ui/react'
import FileList from '../components/ui/FileList'
import UploadedScreenshots from '../components/ui/UploadedScreenshots'
import Form from '../components/ui/Form'
import Hero from '../components/sections/Hero'

import React, { useState } from 'react'

const Create = () => {
  const [isUploaded, setIsUploaded] = useState(false)
  const [fileData, setFileData] = useState({})

  return (
    <>
      <Flex width='full' align='center' justifyContent='center'>
        <Box
          p={8}
          maxWidth='500px'
          borderWidth={1}
          borderRadius={8}
          boxShadow='lg'
        >
          <Box textAlign='center' pb={6}>
            <Heading>Upload Your Screenshots</Heading>
          </Box>
          {isUploaded ? <UploadedScreenshots files={fileData}/> : (
            <FileList 
              uploaded={isUploaded => setIsUploaded(isUploaded)}
              fileData={fileData => setFileData(fileData)}
              canRemove={false}
            />
          )}
          
        </Box>
      </Flex>
    </>
  )
}
export default Create
