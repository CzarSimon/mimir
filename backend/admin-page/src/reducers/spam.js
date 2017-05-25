import { createAction } from 'redux-actions';
import { createHttpObject} from '../methods/helper-methods';
import { baseUrl } from '../config';
import _ from 'lodash';

export const FETCH_SPAM_CANDIDATES = "FETCH_SPAM_CANDIDATES"
export const RECIVE_SPAM_CANDIDATES = "RECIVE_SPAM_CANDIDATES"
export const LABEL_TWEET = "LABEL_TWEET"
export const RECIVE_LABELING_SUCCESS = "RECIVE_LABELING_SUCCESS"
export const RECIVE_LABELING_FAILURE = "RECIVE_LABELING_FAILURE"
export const SKIP_CANDIDATE = "SKIP_CANDIDATE"

const initalState = {
  candidates: [],
  labeledCount: 0
}

const spam = (state = initalState, action = {}) => {
  switch (action.type) {
    case RECIVE_SPAM_CANDIDATES:
      return {
        ...state,
        candidates: action.payload.candidates
      }
    case RECIVE_LABELING_SUCCESS:
      return {
        ...state,
        labeledCount: state.labeledCount + 1,
        candidates: _.tail(state.candidates)
      }
    case RECIVE_LABELING_FAILURE:
      return {
        ...state,
        candidates: _.tail(state.candidates)
      }
    case SKIP_CANDIDATE:
      return {
        ...state,
        candidates: _.tail(state.candidates)
      }
    default:
      return state
  }
}

export default spam;

/* --- Actions --- */
export const reciveSpamCandidates =
  createAction(RECIVE_SPAM_CANDIDATES, candidates => ({candidates}));

export const skipCandidate = createAction(SKIP_CANDIDATE);

export const reciveLabelingSuccess = createAction(RECIVE_LABELING_SUCCESS);

export const reciveLabelingFailure = createAction(RECIVE_LABELING_FAILURE);

export const labelTweet = (labeledTweet, token) => {
  const httpObject = createHttpObject('POST', token, labeledTweet);
  return dispatch => (
    fetch(`${baseUrl}/label-spam`, httpObject)
    .then(res => dispatch(reciveLabelingSuccess()))
    .catch(err => {
      console.log(err);
      dispatch(reciveLabelingFailure())
    })
  )
}

export const fetchSpamCandidates = token => (
  dispatch => (
    fetch(`${baseUrl}/spam-candidates`, createHttpObject('GET', token))
    .then(res => res.json())
    .then(candidates => { dispatch(reciveSpamCandidates(candidates)) })
    .catch(err => { console.log(err) })
  )
)
