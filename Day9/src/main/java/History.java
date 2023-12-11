import java.util.Arrays;

public class History implements Comparable<History> {

  Integer[] values;
  private Integer nextValue;
  private Integer valueBefore;

  public History(Integer[] values) {
    this.values = values;
    this.nextValue = 0;
  }


  @Override
  public int compareTo(History toCompare) {
    return Integer.compare(this.values.length, toCompare.values.length);
  }

  public Integer calculate(Integer base) {
    nextValue = values[values.length - 1] + base;
    return nextValue;
  }

  public Integer calculateReverse(Integer base) {
    valueBefore = values[0] - base;
    return valueBefore;
  }
}
