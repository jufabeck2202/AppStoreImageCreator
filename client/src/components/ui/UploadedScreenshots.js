import {
  Accordion,
} from '@chakra-ui/react'
import ScreenshotItem from '../ui/ScreenshotItem'


const UploadedScreenshots = ({files}) => {
  const fileItems = files.map((file) =>
    <ScreenshotItem file={file}/>
  );

  return (
    <Accordion defaultIndex={[0]} allowMultiple>
      {fileItems}
    </Accordion>
  )
}
export default UploadedScreenshots
