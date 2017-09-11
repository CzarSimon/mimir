import { isEmpty } from '../helper-methods';
import _ from 'lodash';

const BELOW_MEAN_SCORE = 0.0;
const BELOW_1_SIGMA_SCORE = 2.0;
const BELOW_2_SIGMA_SCORE = 4.3;
const ABOVE_2_SIGMA_SCORE = 44.0;

const WEIGHTED_TOTAL = 300.0; // 300.0 Sum of all integer between 1 and 24

const TOTAL_BELOW_MEAN_SCORE = 0.0;
const TOTAL_BELOW_1_SIGMA_SCORE = WEIGHTED_TOTAL * BELOW_1_SIGMA_SCORE;
const TOTAL_BELOW_2_SIGMA_SCORE = WEIGHTED_TOTAL * BELOW_2_SIGMA_SCORE;
const TOTAL_ABOVE_2_SIGMA_SCORE = WEIGHTED_TOTAL * ABOVE_2_SIGMA_SCORE;

export const classifyUrgency = ({ volume, mean, stdev, minute, hour, ticker }) => {
  if (isEmpty(mean) || isEmpty(hour) || isEmpty(stdev) || isEmpty(volume) || isEmpty(minute)) {
    return "low";
  }
  const MINUTES_IN_HOUR = 60;
  volume[hour] = volume[hour] * MINUTES_IN_HOUR / minute;
  const volumeScore = calculateVolumeScore(hour, volume, mean, stdev);
  return urgencyLevel(volumeScore);
}

const urgencyLevel = score => {
  if (score < TOTAL_BELOW_1_SIGMA_SCORE) {
    return "low";
  } else if (score < TOTAL_BELOW_2_SIGMA_SCORE) {
    return "high";
  }
  return "urgent";
}

const calculateVolumeScore = (hour, volumes, mean, stdev) => {
  const scores = mapScore(volumes, mean, stdev);
  const weight = createWeights(hour);
  return _.reduce(scores, (result, score, i) => result + score * weight[i], 0)
}

const mapScore = (volumes, mean, stdev) => (
  _.map(volumes, (v, i) => scoreVolume(mean[i], stdev[i], v))
)

const createWeights = hour => {
  const HOURS_IN_DAY = 24;
  return _.map(Array(HOURS_IN_DAY), (v, i) => (
    (i > hour) ? i - hour : HOURS_IN_DAY - hour + i
  ));
}

const scoreVolume = (μ, σ, v) => {
  if (v < μ) {
    return BELOW_MEAN_SCORE;
  } else if (v < μ + σ) {
    return BELOW_1_SIGMA_SCORE;
  } else if (v < μ + 2 * σ) {
    return BELOW_2_SIGMA_SCORE;
  }
  return ABOVE_2_SIGMA_SCORE;
}
