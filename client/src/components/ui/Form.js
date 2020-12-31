import React from 'react'
import { useForm} from 'react-hook-form'
import {
  RadioGroup,
  Radio,
  HStack,
  Box,
  FormControl,
  Button,
  FormLabel
} from '@chakra-ui/react'

export default function Form () {
  const { handleSubmit, errors, register, formState, watch } = useForm()
  const onSubmit = data => console.log(data)
  const background = watch("background")
  console.log(errors)

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <FormControl>
        
        <Box border='1px'mt={2} mb={2} p={2} borderColor='gray.200' borderRadius={10}>
          <FormLabel>Select Background</FormLabel>
          <RadioGroup name='mode' defaultValue='gradient'>
            <HStack spacing='24px'>
              <Radio
                id='transparent'
                name='background'
                value='transparent'
                ref={register}
              >
                Transparent
              </Radio>
              <Radio
                id='background'
                name='background'
                value='color'
                ref={register}
              >
                Single Color
              </Radio>
              <Radio
                id='gradient'
                name='background'
                value='gradient'
                ref={register}
              >
                Gradient
              </Radio>
            </HStack>
          </RadioGroup>
        </Box>
        <Box border='1px'mt={2} mb={2} p={2} borderColor='gray.200' borderRadius={10}>
        <FormLabel>Select Mode</FormLabel>
        <RadioGroup name='mode' defaultValue='single'>
          <HStack spacing='24px'>
            <Radio id='single' name='mode' value='single' ref={register}>
              Single Frame
            </Radio>
            <Radio id='concat' name='mode' value='concat' ref={register}>
              Merge into one
            </Radio>
          </HStack>
        </RadioGroup>
        </Box>
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
