import React from "react";
import { ChakraProvider } from "@chakra-ui/react";
import {
  Switch,
  Route,
  BrowserRouter
} from "react-router-dom";

import './App.css';
import Landing from "./pages/Landing";
import Create from "./pages/Create";


function App() {
  return (
    <ChakraProvider>
      <BrowserRouter>
      <Switch>
        <Route exact path="/">
          <Landing />
        </Route>
        <Route path="/create">
          <Create />
        </Route>
      </Switch>
    </BrowserRouter>
    </ChakraProvider>
  );
}

export default App;
