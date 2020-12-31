import { Flex, Heading, Box } from '@chakra-ui/react'
import FileList from '../components/ui/FileList'
import UploadedScreenshots from '../components/ui/UploadedScreenshots'


import React, { useState } from 'react'
import useExecute from '../utils/submitHook'

const Create = () => {
  const [isUploaded, setIsUploaded] = useState(false)
  const [id, setID] = useState(false)
  const [fileData, setFileData] = useState({})
  const [execute, response, loading, hasError] = useExecute(`http://localhost:8080/api/process/${id}`)
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

          {isUploaded ? (
            <>
              <UploadedScreenshots files={fileData} submit={execute} />
            </>
          ) : (
            <FileList
              uploaded={isUploaded => setIsUploaded(isUploaded)}
              fileData={fileData => setFileData(fileData)}
              clientID={id => setID(id)}
              canRemove={false}
            />
          )}
        </Box>
      </Flex>
    </>
  )
}
export default Create
