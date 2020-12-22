import { Container, Heading } from '@chakra-ui/react'
import FileList from '../components/ui/FileList'
import ScreenshotPreview from "../components/ui/ScreenshotPreview"
import Form from "../components/ui/Form"

import React, { useState } from 'react'


const Create = () => {
  const [isUploaded, setIsUploaded] = useState(false)
  const [fileData, setFileData] = useState({})
  

  return (
    <Container maxW='xl' centerContent>
      <Heading
        as='h1'
        size='xl'
        fontWeight='bold'
        color='primary.800'
        textAlign={['center', 'center', 'left', 'left']}
      >
        Upload Your Screnshots 
      </Heading>
      <Form/>
      <p>Files Uploaded {isUploaded.toString()}</p>
      <p>Files Data {Object.keys(fileData).length}</p>
      {isUploaded ? <ScreenshotPreview url={fileData[0].meta.previewUrl}/> : (
        <FileList
          uploaded={isUploaded => setIsUploaded(isUploaded)}
          fileData={fileData => setFileData(fileData)}
          canRemove={false}
        />
      )}
    </Container>
  )
}
export default Create
