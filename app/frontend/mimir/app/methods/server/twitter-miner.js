export const urgency_level = (volume, mean, stdev, minute) => {
  const damping = parseFloat(minute) / 60.0;
  if (volume <= (damping * (mean + stdev)) || !volume) {
    return "low";
  } else if (volume <= (damping * (mean + 2 * stdev))) {
    return "high";
  } else {
    return "urgent";
  }
}
