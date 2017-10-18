import { includes } from 'lodash';
import { createAction } from 'redux-actions';
import { getRequest } from '../methods/api-methods';

/* --- Types --- */
export const FETCH_NEWS_ITEMS = 'mimir/news/FETCH';
export const RECIVE_NEWS_FAILURE = 'mimir/news/RECIVE_FAIL';
export const RECIVE_NEWS_ITEMS = 'mimir/news/RECIVE';
export const RECIVE_SHOWCASE_NEWS = 'mimir/news/showcase/RECIVE';
export const SWITCH_PERIOD = 'mimir/news/period/SWITCH';

const initalState = {
  period: 'TODAY',
  defaultTickers: ['AAPL', 'SNAP', 'TSLA', 'AMZN', 'MSFT'],
  showcaseNews: []
}

/* --- Reducer --- */
const news = (state = initalState, action = {}) => {
  switch (action.type) {
    case RECIVE_NEWS_ITEMS:
      return {
        ...state,
        ...action.payload.data
      };
    case RECIVE_SHOWCASE_NEWS:
      return {
        ...state,
        showcaseNews: action.payload.data
      }
    case SWITCH_PERIOD:
      const validPeriods = ['TODAY', '1W', '1M'];
      return (
        (includes(validPeriods, action.payload.period)) ? (
          {
            ...state,
            period: action.payload.period
          }
        ) : (
          state
        )
      );
    default:
      return state;
  }
}
export default news;

/* --- Actions --- */
export const fetchNewsItems = (ticker, period) => (
  dispatch => (
    getRequest(`api/news/${ticker}/5/${period}`)
    .then(news => dispatch(reciveNewsItems({ [ticker]: news })))
    .catch(err => {
      console.log(err);
    })
  )
);

export const reciveNewsItems = createAction(
  RECIVE_NEWS_ITEMS, data => ({ data })
);

export const reciveShowcaseNews = createAction(
  RECIVE_SHOWCASE_NEWS, data => ({ data })
);

export const switchPeriod = createAction(
  SWITCH_PERIOD, period => ({ period })
);
