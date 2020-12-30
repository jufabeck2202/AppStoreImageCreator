
import { HexColorPicker } from "react-colorful";
import { useState} from 'react'

import "react-colorful/dist/index.css";

const Picker = () => {
  const [color, setColor] = useState("#aabbcc");
  return <HexColorPicker color={color} onChange={setColor} />;
};

export default Picker;