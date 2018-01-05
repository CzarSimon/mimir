import React, { Component } from 'react'
import Description from '../../components/util/description'
import ButtonControlsContainer from '../containers/button-controls'
import DescriptionEditContainer from '../containers/description-edit'
import { length, color } from '../../styles/styles'

const styles = {
  fullInfo: {
    marginTop: length.small
  },
  siteLink: {
    marginBottom: length.small
  },
  link: {
    color: color.blue,
    textDecoration: 'none'
  }
}

export default class FullInfo extends Component {
  render() {
    const { Ticker, Description: desc, Website, editMode } = this.props
    const descriptionComponent = (!editMode)
    ? <Description text={desc} />
    : <DescriptionEditContainer currentDescription={desc} Ticker={Ticker}/>
    return (
      <div style={styles.fullInfo}>
        {(Website) ? (
          <p style={styles.siteLink}>
            Website: <a style={styles.link} target='_blank' href={Website}>{Website}</a>
          </p>
        ) : <div />}
        { descriptionComponent }
        <ButtonControlsContainer Ticker={Ticker} editMode={editMode} />
      </div>
    )
  }
}
