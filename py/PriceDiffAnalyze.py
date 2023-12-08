import pandas as pd

# Prices in NT Dollars for iPhone 15 Pro models
prices = {
    128: 36900,
    256: 40400,
    512: 47400,
    1024: 54400  # 1TB is 1024GB
}

# Initialize a DataFrame from the prices dictionary
data_df = pd.DataFrame(list(prices.items()), columns=["GB", "Price"])
# Calculate and add a column for price per GB
data_df["Price Per GB"] = data_df["Price"] / data_df["GB"]

# Initialize a DataFrame for price differences
price_diff_df = pd.DataFrame()

# Calculate additional price and GB for each upgrade
price_diff_df["Add Price"] = data_df["Price"].diff().shift(-1)
price_diff_df["Add GB"] = data_df["GB"].diff().shift(-1)

# Drop the last row from both dataframes as it does not have a valid comparison
data_df = data_df[:-1]
price_diff_df = price_diff_df[:-1]

# Calculate Price Difference Per Additional GB and round to 2 decimal places
price_diff_df["Price Diff Per Add GB"] = (price_diff_df["Add Price"] / price_diff_df["Add GB"]).round(2)

# Convert Add GB to integer
price_diff_df["Add GB"] = price_diff_df["Add GB"].astype(int)

# Adding 'spec' column for storage upgrade specification
price_diff_df["GB Diff"] = \
    data_df["GB"].astype(str) + " to " + (data_df["GB"] + price_diff_df["Add GB"]).astype(str)

# Reordering columns for clarity
price_diff_df = price_diff_df[["GB Diff", "Add Price", "Add GB", "Price Diff Per Add GB"]]

# Output the DataFrames
print("Data DataFrame:")
print(data_df)
print("\nPrice Difference DataFrame:")
print(price_diff_df)
