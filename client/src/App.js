import React from 'react'
import {
  ChakraProvider,
  ThemeProvider,
  theme,
  CSSReset
} from '@chakra-ui/react'
import { ColorModeProvider } from "@chakra-ui/color-mode"

import { Switch, Route, BrowserRouter } from 'react-router-dom'

import './App.css'
import Landing from './pages/Landing'
import Create from './pages/Create'
import ThemeToggler from './utils/ThemeToggler'

function App () {
  return (
    <ChakraProvider>
      <ThemeProvider theme={theme}>
      <ColorModeProvider options={{
        useSystsemColorMode: true
      }}>
          <CSSReset />
          <ThemeToggler />
          <BrowserRouter>
            <Switch>
              <Route exact path='/'>
                <Landing />
              </Route>
              <Route path='/create'>
                <Create />
              </Route>
            </Switch>
          </BrowserRouter>
        </ColorModeProvider>
      </ThemeProvider>
    </ChakraProvider>
  )
}

export default App
