'use strict';
import * as types from './action-types';
import { createAction } from 'redux-actions';
import { query_knowledge_graph } from '../methods/google-api';

export const fetch_company_desc = (company_name, ticker) => {
  return (dispatch) => {
    return query_knowledge_graph(company_name)
    .then(desc => dispatch(recive_company_desc(desc, ticker)))
    .catch(err => dispatch(recive_desc_failure(err, ticker)))
  }
}

export const recive_desc_failure = createAction(types.RECIVE_DESC_FAILURE,
  (error, ticker) => ({ error, ticker })
)

export const recive_company_desc = createAction(types.RECIVE_COMPANY_DESC,
  (description, ticker) => ({ description, ticker })
)
