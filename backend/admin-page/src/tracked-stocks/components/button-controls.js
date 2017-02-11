import React, { Component } from 'react';
import Button from '../../components/util/button';
import { length, color, font } from '../../styles/styles';

export default class ButtonControlls extends Component {
  render() {
    return (
      <div>
        <Button text={Edit}/>
        <Button
          text={Untrack stock}
          customStyles={{backgroundColor: color.red}}
        />
      </div>
    )
  }
}
