import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

public class Oasis {

  private List<History> histories = new ArrayList<>();

  public History buildHistory(Integer[] values) {
    History history = new History(values);
    histories.add(history);
    return history;
  }

  public void extrapolateToZeroes() {
    Collections.sort(histories);
    var values = histories.get(0).values;
    boolean allZeroed = false;
    while (!allZeroed) {
      allZeroed = true;
      Integer[] nextValues = new Integer[values.length -1];
      for (int i = 0; i < values.length - 1; i++) {
        nextValues[i] = values[i + 1] - values[i];
        if (nextValues[i] != 0) {
          allZeroed = false;
        }
      }
      values = nextValues;
      buildHistory(nextValues);
    }
  }

  public Integer calculateNextValue() {
    Collections.sort(histories);
    Integer partialValue = 0;
    for (History history : histories) {
      partialValue = history.calculate(partialValue);
    }
    return partialValue;
  }

  public Integer calculateNextValueReversed() {
    Collections.sort(histories);
    Integer partialValue = 0;
    for (History history : histories) {
      partialValue = history.calculateReverse(partialValue);
    }
    return partialValue;
  }
}
