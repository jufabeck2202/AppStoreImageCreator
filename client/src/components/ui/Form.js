import React, { useEffect } from 'react'
import { useForm } from 'react-hook-form'
import {
  RadioGroup,
  Radio,
  HStack,
  Box,
  FormControl,
  Button,
  FormLabel,
  Center
} from '@chakra-ui/react'
import Picker from './Picker'

export default function Form () {
  const {
    handleSubmit,
    errors,
    register,
    formState,
    watch,
    setValue
  } = useForm()
  const onSubmit = data => console.log(data)
  const background = watch('background')

  useEffect(() => {
    setValue('background', 'gradient')
  }, [])

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <FormControl>
        <Box
          border='1px'
          mt={2}
          mb={2}
          p={2}
          borderColor='gray.200'
          borderRadius={10}
        >
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
          <Center>
            <Box p={2}>
              {background === 'color' && (
                <Picker
                  text='Select Color'
                  name='color'
                  register={register}
                  setValue={setValue}
                />
              )}
              {background === 'gradient' && (
                <>
                    <Picker
                      text='Select Color 1'
                      name='gradient1'
                      register={register}
                      setValue={setValue}
                    />
                    <Picker
                      text='Select Color 2'
                      name='gradient2'
                      register={register}
                      setValue={setValue}
                    />
                </>
              )}
            </Box>
          </Center>
        </Box>
        <Box
          border='1px'
          mt={2}
          mb={2}
          p={2}
          borderColor='gray.200'
          borderRadius={10}
        >
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
