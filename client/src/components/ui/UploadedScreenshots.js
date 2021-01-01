import { Accordion, Button } from '@chakra-ui/react'
import ScreenshotItem from '../ui/ScreenshotItem'
import BottomForm from '../ui/Form'



const UploadedScreenshots = ({ files, submit }) => {
  const fileItems = files.map(file => <ScreenshotItem file={file} />)

  return (
    <>
      <Accordion defaultIndex={[0]} allowMultiple>
        {fileItems}
      </Accordion>
      <BottomForm />
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
