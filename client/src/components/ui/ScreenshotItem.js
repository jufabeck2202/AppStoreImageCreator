import {
    Accordion,
    AccordionItem,
    AccordionButton,
    AccordionPanel,
    AccordionIcon,
    Box,
    Badge,
    Image,
    HStack,
    Text,
    Input
  } from '@chakra-ui/react'
  
  const ScreenshotItem = ({file}) => {
    return (
        <AccordionItem>
          <AccordionButton>
            <Box flex='1' textAlign='left'>
              {file.file.name}
            </Box>
            <Badge mr={2} colorScheme='green'>{file.meta.device}</Badge>
            <Badge colorScheme='blue'>has Text</Badge>
            <AccordionIcon />
          </AccordionButton>
          <AccordionPanel pb={4}>
          <HStack spacing={2}>
            <Box  boxSize="150px">
              <Image src={file.meta.previewUrl} alt='Segun Adebayo' />
            </Box>
            <Box  w="100%" pb={10}>
            <Text fontSize="md">{file.file.name}</Text>
            <Input placeholder="Add Heading"/>
            </Box>
            </HStack>
          </AccordionPanel>
        </AccordionItem>
    )
  }
  export default ScreenshotItem
  