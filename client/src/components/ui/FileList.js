import 'react-dropzone-uploader/dist/styles.css'
import Dropzone from 'react-dropzone-uploader'
import { useState } from 'react'

const FileUploader = ({ uploaded, fileData }) => {
  const [id, setID] = useState(0)

  // specify upload params and url for your files
  const getUploadParams = ({ meta }) => {
    console.log('upload')
    const idReturn = '/' + id
    return {
      url: `http://localhost:8080/api/upload${id !== 0 ? idReturn : ''}`,
      method: 'POST'
    }
  }

  // called every time a file's `status` changes
  const handleChangeStatus = ({ meta, file, xhr }, status) => {
    console.log('status change')
    console.log(status, meta, file)
    if (status === 'done') {
      var json = JSON.parse(xhr.response)
      setID(json.id)
    }
  }

  // receives array of files that are done uploading when submit button is clicked
  const handleSubmit = (files, allFiles) => {
    console.log(files)
    console.log(allFiles)
    uploaded(true)
    fileData(allFiles)
  }

  return (
    <Dropzone
      getUploadParams={getUploadParams}
      onChangeStatus={handleChangeStatus}
      onSubmit={handleSubmit}
      accept='image/*'
      canCancel={false}
      canRemove={false}
      canRestart={false}
    />
  )
}
export default FileUploader
