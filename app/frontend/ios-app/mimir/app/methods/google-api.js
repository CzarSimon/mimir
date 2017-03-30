'use strict'
import { knowledge_graph } from '../credentials/google';

export const kgSearch = name => {
  const api = "https://kgsearch.googleapis.com/v1/entities:search";
  const queryURL = api + '?query=' + name + '&key=' + knowledge_graph.key + '&limit=1&indent=True';
  return (
    fetch(queryURL)
    .then((res) => res.json())
    .then((res) => res.itemListElement[0].result.detailedDescription.articleBody)
  );
}
