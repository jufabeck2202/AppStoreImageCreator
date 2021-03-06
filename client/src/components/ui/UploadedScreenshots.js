import { Accordion, Heading, Box } from '@chakra-ui/react'
import ScreenshotItem from '../ui/ScreenshotItem'
import BottomForm from '../ui/Form'

const UploadedScreenshots = ({ files, submit, isLoading }) => {
  const fileItems = files.map(file => <ScreenshotItem file={file} />)

  return (
    <>
      <Box textAlign='center' pb={6}>
        <Heading>Apply your Effects</Heading>
      </Box>
      <Accordion defaultIndex={[0]} allowMultiple>
        {fileItems}
      </Accordion>
      <BottomForm submit={submit} isLoading={isLoading} />
    </>
  )
}
export default UploadedScreenshots
