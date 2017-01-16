import _ from 'lodash';

export const portraitMode = () => {
  const { innerHeight, innerWidth } = window
  return innerHeight > innerWidth
}
