import 'react-dropzone-uploader/dist/styles.css'
import Dropzone from 'react-dropzone-uploader'
import { useState } from 'react'

const FileUploader = ({ uploaded, fileData, clientID }) => {
  const [id, setID] = useState(0)

  // specify upload params and url for your files
  const getUploadParams = ({ file, meta }) => {
    const idReturn = '/' + id
    const body = new FormData()
    body.append('file', file)
    body.append('height',meta.height)
    body.append('width',meta.width)

    return {
      url: `http://localhost:8080/api/upload${id !== 0 ? idReturn : ''}`,
      body
    }
  }

  // called every time a file's `status` changes
  const handleChangeStatus = ({ meta, file, xhr }, status) => {
    if (status === 'done') {
      var json = JSON.parse(xhr.response)
      setID(json.id)
      clientID(json.id)
      meta["device"] = json.device
    }
  }

  // receives array of files that are done uploading when submit button is clicked
  const handleSubmit = (files, allFiles) => {
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
