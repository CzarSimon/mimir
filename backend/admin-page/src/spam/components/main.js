import React, { Component } from 'react'
import MainMenu from '../../main-menu/components/main-menu';
import Loading from '../../components/util/loading';
import PageTitle from '../../components/util/page-title';
import CandidateCard from '../components/candidate-card';


export default class Spam extends Component {
  render() {
    const component = (this.props.loaded)
    ? <CandidateCard {...this.props} />
    : <Loading />
    return (
      <div className='fullpage'>
        <MainMenu />
        <div className='content'>
          <PageTitle title={"Label spam"} />
          {component}
        </div>
      </div>
    )
  }
}
