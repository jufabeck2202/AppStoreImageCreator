import React from 'react'
import { useForm } from 'react-hook-form'
import { Select, Input, Checkbox, FormControl, Button,} from '@chakra-ui/react'

export default function Form () {
  const { handleSubmit, errors, register, formState } = useForm()
  const onSubmit = data => console.log(data)
  console.log(errors)

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <FormControl isInvalid={errors.name}>
        <Input
          type='text'
          placeholder='Heading'
          name='Heading'
          ref={register}
        />

        <Checkbox
          type='checkbox'
          placeholder='LargeImage'
          name='LargeImage'
          ref={register} 
        > Genrate Color 
        </Checkbox>
        <Select name='SelectMode' ref={register}>
          <option value='concat'>concat</option>
          <option value=' Preview'> Preview</option>
          <option value=' AppStore'> AppStore</option>
        </Select>

        <Button
          mt={4}
          colorScheme='teal'
          isLoading={formState.isSubmitting}
          type='submit'
        >
          Submit
        </Button>
      </FormControl>
    </form>
  )
}
