import java.util.Arrays;
import java.util.List;
import org.assertj.core.api.Assertions;
import org.junit.jupiter.api.Test;

public class Day9Test {

  @Test
  void test__givenInput__calculateNextValue() {
    var given = "0 3 6 9 12 15";
    Oasis oasis = new Oasis();
    Integer[] values = Arrays.stream(given.split(" ")).map(Integer::parseInt)
        .toArray(Integer[]::new);
    oasis.buildHistory(values);
    oasis.extrapolateToZeroes();
    Integer nextValue = oasis.calculateNextValue();
    Assertions.assertThat(nextValue).isEqualTo(18);
  }

  @Test
  void test__givenMultipleInput__calculateSumOfNextValue() {
    var given = List.of("0 3 6 9 12 15", "1 3 6 10 15 21", "10 13 16 21 30 45");
    Day9 day9 = new Day9();
    Integer total = day9.calculateAllEntries(given);
    Assertions.assertThat(total).isEqualTo(114);
  }
}
