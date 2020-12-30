import { Accordion, Button } from '@chakra-ui/react'
import ScreenshotItem from '../ui/ScreenshotItem'
import Picker from './Picker'


const UploadedScreenshots = ({ files, submit }) => {
  const fileItems = files.map(file => <ScreenshotItem file={file} />)

  return (
    <>
      <Accordion defaultIndex={[0]} allowMultiple>
        {fileItems}
      </Accordion>
      <Picker/>
      <Button
        colorScheme='teal'
        variant='outline'
        onClick = {submit}
      >
        Submit
      </Button>
    </>
  )
}
export default UploadedScreenshots
